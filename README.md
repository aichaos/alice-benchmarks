# Alice Benchmarks

This repository is for testing how well an [A.L.I.C.E.][1]-sized brain performs
on the various implementations of RiveScript. The goal is to have a benchmark
result set showing how slow the implementations are at handling a brain of this
size, which can also be used to compare improvements that make Alice run better
in RiveScript.

Contributions are welcome. If one of these benchmark scripts turns out not to be
doing a very good job of benchmarking, send a pull request that fixes it.

This repository will also be used to discuss ideas to improve the performance of
RiveScript. See the [Issues][7] for that.

## Setup

You'll need to install the various RiveScript modules for each language before
running their scripts:

* Go: `go get github.com/aichaos/rivescript-go`
* JavaScript: `npm install` (installs RiveScript and other dependencies)
* Perl: `cpanm RiveScript`
* Python: `pip install rivescript`

# Results

* [Results-20160310](./Results-20160310.md) -- initial testing done on
  March 10 2016.

# Misc Notes

There are some weird quirks with A.L.I.C.E.

The RiveScript files were downloaded from [here][6] where they were
originally converted on Feb 13 2013. A handful of triggers had empty
responses (just a `-` symbol with no data).

An interesting quirk was that the "hello" trigger would redirect to "hi"
about 50% of the time, but there is no trigger that matches "hi". This
would cause certain implementations of RiveScript to run away and get a
deep recursion error (in particular the JavaScript implementation). The
recursive chain ends up on the `push random` trigger (`push *` in stack.rive).

Some replies end up redirecting to "random pickup line", but there is no trigger
for this. It ends up matching the `*` trigger, which goes back to
"random pickup line" sometimes and rinse & repeat. So I removed all references
to "random pickup line" from the code.

The message "Hello Alice" kicks off a series of different redirects and it makes
this reply take the longest. First it has to scan all through the sorted replies
until it reaches `hello *`, which redirects to `hello` + redirects to `Alice`.
There is no trigger of simply "Alice", so it ends up reaching the `*` trigger
which redirects to `random pickup line` in addition to redirecting to `push *`.

At some point in the future I'll probably put the RiveScript fork of A.L.I.C.E.
on GitHub and begin cleaning up its strange quirks. As it currently stands, it
poses a real risk of bringing your chatbot down if the wrong chain of triggers
gets activated.

# License

The Alice reply set here is based on the [Free A.L.I.C.E. AIML Set][2],
converted to RiveScript using the [aiml2rs][3] script.

The Free A.L.I.C.E. AIML set is released by Dr. Richard Wallace under the
terms of the [GNU Lesser GPL][4] license.

The RiveScript modules used in this project are released under the
[MIT License][5].

The programming language source codes in *this* repository are released as
public domain code.

[1]: https://en.wikipedia.org/wiki/Artificial_Linguistic_Internet_Computer_Entity
[2]: https://code.google.com/archive/p/aiml-en-us-foundation-alice/
[3]: https://github.com/aichaos/aiml2rs
[4]: http://www.gnu.org/licenses/lgpl.html
[5]: https://opensource.org/licenses/MIT
[6]: https://static.rivescript.com/files/sets/AliceRS-0.03.tar.gz
[7]: https://github.com/aichaos/alice-benchmarks/issues
