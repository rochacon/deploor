#!/usr/bin/env python
import os
import sys
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..', '..', 'clone')));

from fabric.colors import cyan, red
from fabfile import deploy


# Override os.environ['ENVIRONMENT'] with $PWD info
environ = os.getcwd().split('/')[-2]
if environ in ('dev', 'staging', 'production'):
    os.environ['ENVIRONMENT'] =  environ
else:
    print red('Unknown environment')
    sys.exit(2)


def parse_git_ref(git_ref):
    ref = git_ref.split('/')
    if len(ref) < 3:
        return git_ref

    # Production only deploys tags
    if os.environ.get('ENVIRONMENT', 'unknown') =='production' and ref[1] != 'tags':
        print red('Only tags can be deployed to production')
        sys.exit(1)

    # Otherwise, return the branch name
    return ref[2]


deploy(parse_git_ref(sys.argv[1]))
