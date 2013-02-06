# -*- coding: utf-8 -*-
import os
import shutil
import sys

from fabric.colors import red


def abort(message, code, clone=None):
    print red(message)
    if clone is not None:
        shutil.rmtree(clone)
    sys.exit(code)


def parse_git_ref(git_ref):
    """
    Return the reference type and name
    Example:
        refs/head/master -> branch, master
        refs/tags/1.0 -> tags, 1.0
    """
    ref = git_ref.split('/')
    if len(ref) < 3:
        raise ValueError('invalid ref string')

    # Otherwise, return the branch name
    ref_type = 'branch' if ref[1] == 'head' else ref[1]
    return ref_type, ref[2]


def get_enviroment():
    '''
    Based on a repository path, get the environment

    Example: /srv/git/staging/project.git
    '''
    return os.getcwd().split('/')[-2]

