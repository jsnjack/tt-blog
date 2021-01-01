---
title: "Test context of the rendered template with RequestFactory in Django"
date: 2017-03-01T23:18:13+01:00
draft: false
tags: ["linux", "python", "django"]
---

When writing tests for the Django's view functions I prefer to use RequestFactory instead of the django test client. Tests that use RequestFactory are easier to understand and they are faster. For example, you need to test a simple view function:
```python
def incoming_queue(request):
    if request.method == "POST":
        pass
    else:
        incoming_queue_list = UsageHistory.objects.filter(company=request.user.agent.company, queued=True)
        context = {
            "incoming_queue_list": incoming_queue_list,
            "page": "queue",
            "agent_available_url": reverse("agent_available")
        }
        return render(request, 'agents/incoming_queue.html', context)
```

If you want to test the context of the rendered template, it is easy to do with django test client as it attaches it to the response. If you want to use RequestFactory for that purpose, there is no easy way to do it. It would be comfortable to have something like assertTemplateContextIn(expected_context) method, that checks that expected_context was used during the template rendering. The following code was inspired by assertTemplateUsed method:
```python
from copy import copy

from django.db.models.query import QuerySet
from django.test.signals import template_rendered
from django.test.utils import ContextList


class TemplateContextTestCase(object):
    def assertTemplateContextIn(self, expected_context=None):
        """
        Asserts that the template was rendered with the context that includes
        expected_context. Example:
            with self.assertTemplateContextIn(expected_context):
                vuew_function(self.request)
        """
        return _AssertTemplateContext(self, expected_context)


class _AssertTemplateContext(object):
    def __init__(self, test_case, expected_context):
        self.test_case = test_case
        self.expected_context = expected_context
        self.context = ContextList()

    def on_template_render(self, sender, signal, template, context, **kwargs):
        self.context.append(copy(context))

    def test(self):
        missing = []
        mismatched = []
        for key in self.expected_context.keys():
            if key in self.context:
                if self.context[key] != self.expected_context[key]:
                    match = False
                    # Compare QuerySets
                    if isinstance(self.context[key], QuerySet) and isinstance(self.expected_context[key], QuerySet):
                        if self.context[key].model == self.expected_context[key].model:
                            if list(self.context[key].values_list("id", flat=True)) == \
                               list(self.expected_context[key].values_list("id", flat=True)):
                                match = True
                    if not match:
                        mismatched.append(key)
            else:
                missing.append(key)
        return missing, mismatched

    def __enter__(self):
        template_rendered.connect(self.on_template_render)
        return self

    def __exit__(self, exc_type, exc_value, traceback):
        template_rendered.disconnect(self.on_template_render)
        if exc_type is not None:
            return

        missing, mismatched = self.test()
        message = ""
        if missing:
            message = message + "Missing keys: " + ", ".join(x for x in missing) + "\n"
        if mismatched:
            for item in mismatched:
                message = message + "Context key mismatch '%s': %s != %s \n" % (
                    item, self.expected_context[item], self.context[item]
                )
        if message:
            self.test_case.fail(message)
```

We often pass querysets instances to the context of the rendered template. The code above checks if 2 QuerySets are the same by comparing their models and ids of the elements. TemplateContextTestCase class can be used to extend TestCase class:
```python
from django.test import TestCase, RequestFactory

from yourapp.tests.helpers import TemplateContextTestCase
from django.core.urlresolvers import reverse


class TestIncomingQueue(TestCase, TemplateContextTestCase):
    def test_context(self):
        agent = AgentFactory()
        self.request.user = agent.user
        expected = {
            "page": "queue",
            "agent_available_url": reverse("agent_available"),
            "incoming_queue_list": UsageHistory.objects.none()
        }
        with self.assertTemplateContextIn(expected):
            incoming_queue(self.request)
```

