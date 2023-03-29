// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package esg is the emergent stochastic generator, where tokens are generated
stochastically according to rules defining the contingencies and probabilities.
It can be used for generating sentences (sg as well).

Rules:

There are two types of rules:
* unconditional random items
* conditional items.

Unconditional items are chosen at random, optionally with specified probabilities:

RuleName {
    %50 Rule2 Rule4
    %30 'token1' 'token2'
    ...
}

where Items on separate lines within each rule are orthogonal options,
chosen at uniform random unless otherwise specified with a leading %pct.
%pct can add up to < 100 in which case *nothing* is an alternative output.

Multiple elements in an item (on the same line) are resolved and
emitted in order.
Terminal literals are 'quoted' -- otherwise refers to another rule:
error message will flag missing rules.

Conditional items are specified by the ? after the rule name:

RuleName ? {
    Rule2 || Rule3 {
        Item1
        Item2
        ...
    }
    Rule5 && Rule6 {
        ...
    }
    ...
}

The expression before the opening bracket for each item is a standard logical
expression using || (or), && (and), and ! (not), along with parens,
where the elements are rules that could have been generated earlier in the pass --
they evaluate to true if so, and false if not.

If the whole expression evaluates to true, then it is among items chosen at random
(typically only one for conditionals but could be any number).

If just one item per rule it can be put all on one line.

States:

Each Rule or Item can have an optional State expression associated with it,
which will update a `States` map in the overall Rules if that Rule or Item fires.
The `States` map is a simple `map[string]string` and is reset for every Gen pass.
It is useful for annotating the underlying state of the world implied by the
generated sentence, using various role-filler values, such as those given by
the modifiers below.  Although the Fired map can be used to recover this
information, the States map and expressions can be designed to make a clean,
direct state map with no ambiguity.

In the rules text file, an `=` prefix indicates a state-setting expression
-- it can either be a full expression or get the value automatically:
* `=Name=Value` -- directly sets state Name to Value
* `=Name`  -- sets state Name to value of Item or Rule that it is associated with.
  Only non-conditional Items can be used to set the value, which is the first
  element in the item expression -- conditionals with sub-rules must set the value explicitly.

Expressions at the start of a rule (or sub-rule), on a separate line, are
associated with the rule and activate when that rule is fired (and the implicit
value is the name of the rule).  Expressions at the end of an Item line are
associated with the Item.  Put single-line sub-rule state expressions at end
just before }.

Std Modifiers:

Conventional modifiers, used for defining sub-rules:
A = Agent
	Ao = Co-Agent
V = Verb
P = Patient,
	Pi = Instrument
	Pc = Co-Patient
L = Location
R = adverb
*/
package esg
