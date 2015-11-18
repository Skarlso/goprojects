from multiprocessing import Process
import os
import resource
import subprocess

def run_daemon(command, detach=True):
    pid = os.fork()
    if pid == 0:
      if detach:
        os.setsid()
    else:
      return pid

    # I'm using system here so that this command returns.
    # Could have use Popen or subprocess call which would have been better
    # but the point is to use os.fork()
    # os.execv(command, arg)
    os.system(command)

def run_commands():
    # Couldn't use {1..10} here...
    for x in range (50):
        p = Process(target=run_daemon, args=("for i in `seq 1 10`; do echo Hello; done",))
        p.start()
        p.join() # this blocks until the process terminates

def test_run_commands(benchmark):
    benchmark.pedantic(run_commands, iterations=10, rounds=30)
