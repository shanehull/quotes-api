package main

type Quotes struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quotes{
	{
		Text:   "With all things being equal, the simplest explanation tends to be the right one.",
		Author: "William of Ockham",
	},
	{
		Text:   "The first principle is that you must not fool yourself - and you are the easiest person to fool.",
		Author: "Richard P. Feynman",
	},
	{
		Text:   "There will be no order, only chaos.",
		Author: "Sol Robeson",
	},
	{
		Text:   "There are two kinds of forecasters: those who don't know, and those who don't know they don't know.",
		Author: "John Kenneth Galbraith",
	},
	{
		Text:   "No matter how sophisticated our choices, how good we are at dominating the odds, randomness will have the last word.",
		Author: "Nassim Nicholas Taleb",
	},
	{
		Text:   "The idea is that flowing water never goes stale, so just keep on flowing.",
		Author: "Bruce Lee",
	},
	{
		Text:   "The idea that the future is unpredictable is undermined every day by the ease with which the past is explained.",
		Author: "Daniel Kahneman",
	},
	{
		Text:   "No man ever steps in the same river twice, for it's not the same river and he's not the same man.",
		Author: "Heraclitus",
	},
	{
		Text:   "I am the wisest man alive, for I know one thing, and that is that I know nothing.",
		Author: "Socrates",
	},
	{
		Text:   "If the only tool you have is a hammer, you tend to see every problem as a nail.",
		Author: "Abraham Maslow",
	},
	{
		Text:   "I think that a life properly lived is just learn, learn, learn all the time.",
		Author: "Charles T. Munger",
	},
	{
		Text:   "History never repeats itself. Man always does.",
		Author: "M. de Voltaire",
	},
	{
		Text:   "Do not seek to follow in the footsteps of the wise. Seek what they sought.",
		Author: "Matsuo Basho",
	},
	{
		Text:   "Risk is what's left over after you think you've thought of everything.",
		Author: "Carl Richards",
	},
	{
		Text:   "The beginning of wisdom is the ability to call things by their right names.",
		Author: "Kongzi",
	},
	{
		Text:   "The flame that burns twice as bright burns half as long.",
		Author: "Laozi",
	},
}