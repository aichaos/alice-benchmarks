#!/usr/bin/perl

# Perl RiveScript Benchmark for A.L.I.C.E.

use strict;
use warnings;
use RiveScript;
use Benchmark qw(:all);
use Time::HiRes qw(time);

my @test_inputs = (
	"Hello.",
	"Hello Alice",
	"Who is dr. wallace?",
	"What is your name?",
	"What is my name?",
	"What is the capital of Texas?",
);

# Run all of this 5 times.
timethis(5, sub {
	my $t = time();

	# Time how long it takes to load A.L.I.C.E. from disk.
	my $bot = RiveScript->new();
	$bot->loadDirectory("./alice");

	$t = delta($t, "Loading A.L.I.C.E. from disk.");

	# Time how long it takes to sort the replies.
	$bot->sortReplies();
	$t = delta($t, "Sort the replies.");

	# Get some responses.
	foreach my $input (@test_inputs) {
		my $reply = $bot->reply("user", $input);
		$t = delta($t, "Reply to $input => $reply");
	}

	print "\n";
});

# Helper function that prints the time delta, a message, and returns a new time.
sub delta {
	my ($t, $desc) = @_;
	my $delta = time() - $t;
	printf "[%08fs] %s\n", $delta, $desc;
	return time();
}
