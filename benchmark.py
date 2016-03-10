#!/usr/bin/env python3

# Python RiveScript benchmark for A.L.I.C.E.

import statistics
import time
from rivescript import RiveScript

test_inputs = [
    "Hello.",
    "Hello Alice",
    "Who is dr. wallace?",
    "What is your name?",
    "What is my name?",
    "What is the capital of Texas?",
]

def main():
    """Main entry point."""
    count = timereps(5, test_rs)
    print("Average time of run: {:.8f}s".format(count))

def test_rs():
    """Test how fast RiveScript runs."""
    t = time.time()

    # Load Alice from disk.
    bot = RiveScript()
    bot.load_directory("./alice")
    t = delta(t, "Loading A.L.I.C.E. from disk.")

    # Sort the replies.
    bot.sort_replies()
    t = delta(t, "Sort the replies.")

    # Get some responses.
    for message in test_inputs:
        reply = bot.reply("user", message)
        t = delta(t, "Reply to {} => {}".format(message, reply))

    print("")

def delta(t, desc):
    """Helper function to print a time delta with a message."""
    delta = time.time() - t
    print("[{:.8f}s] {}".format(delta, desc))
    return time.time()

def timereps(reps, func):
    """Helper function to call a function multiple times for benchmarking."""
    times = []
    for i in range(0, reps):
        start = time.time()
        func()
        times.append(time.time() - start)
    return statistics.mean(times)

if __name__ == "__main__":
    main()
