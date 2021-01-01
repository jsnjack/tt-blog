---
title: "python: Django-like testrunner "
date: 2018-03-01T23:27:15+01:00
draft: false
tags: ["linux", "python", "django"]
---

Django-like testrunner for non-django projects:
```python
#!/usr/bin/env python

import argparse
import collections
import unittest
import sys


def main():
    parser = argparse.ArgumentParser(description='Django-like testrunner')
    parser.add_argument('specific_test', metavar='', type=str, nargs='?', help='Specifc test')
    parser.add_argument("--failfast", action="store_true")
    parser.add_argument("--verbosity", type=int, default=1)

    args = parser.parse_args()

    loader = unittest.TestLoader()
    all_tests = loader.discover('.', top_level_dir="./")
    suite = unittest.TestSuite()

    if args.specific_test:

        def walk_tests(tests):
            if isinstance(tests, collections.Iterable):
                for item in tests:
                    walk_tests(item)
                return
            if tests.id().startswith(args.specific_test):
                suite.addTest(tests)
            elif not str(tests).startswith("test"):
                sys.exit("Error in file %s" % tests)

        walk_tests(all_tests)
    else:
        suite.addTests(all_tests)

    unittest.TextTestRunner(verbosity=args.verbosity, failfast=args.failfast).run(suite)


if __name__ == "__main__":
    main()
```
