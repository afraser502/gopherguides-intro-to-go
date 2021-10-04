# History in Technology

I have been working in IT for almost 30 years, which is kind of scary even as I find myself writing this! I learned several scripting / programming languages over the years, vbscript, powershell, bash and Python and also used C# for several years. I've learned some parts of Go however there are plenty of areas that I need to improve upon to further my understanding and fill in the gaps. As is evident from the glut of tooling in the Cloud Native space whicxh is written in Go or has at least strong support for Go tooling, it is a language that will serve well. While I do not have to write day to tooling for my role, it is something that does provide great value not only in regards to being able to provide extensibility, but to be able to simply read Go and understand why parts of the code function the way they do. Often documentation does not define exactly how a piece of code behaves whereas being able to delve into the code and read it yourself is a good way to understand how a piece of software works. As Go is relatively easy to read (in comparisaon to several other languages), it makes for an easier transition.

I think Go will continue to gain traction and support from the community as a whole due to the amount of tooling that is currently available. While products such as Kubernetes, Terraform and Docker to name a few, are growing in adoption at an incredible rate, the proliferation of the ecosystem of tools will make it difficult to manage therefore having a consistent language that developers and engineers can utilise is key towards integrating them together. Go has often been touted as the next best thing although it surprisingly rates in the lower reaches of the top 10 programming languages. I suspect however this is due to comfort with a current language and the time and energy to learn a new language may keep those numbers lower. I personally think in the next 5 to 10 years however that Go will sit much higher in the list of developer languages as time progresses.

# Notes on submission of assignment

The variables assigned uses the short method which assigns the variable and uses Go's automatic inference of the data type. I've used 3 short (terse) variable names as they do not require explanation and will not be used outside of this block of code. I've also chosen to use mutliple assignment for the variables on the same line. The code is quite simple so it saves lines of code without being unreadable.

fmt.Printf is useful in the code submitted this week to display the type of value by simply passing "%T" followed by text "!" followed by a new line "\n" seperator.

The disadvantage to using Printf is that you need to lay out the spacing and line breaks yourself. Using fmt.Println automatically adds spaces between values and appends a new line onto the end of the output.

The go fmt package is also useful in auto formatting your code. While you obviously try to write idiomatic code, sometimes it can be subjective while other times is simply extra or missing spaces. While formatting code is not exclusive to Go, it is very flexible. In Ruby for example, Rubocop will provide linting and formatting whereas Python utilises tools such as black or autopep8. Python is quite annoying in that the Pythonic style guide is quite rigid and can be difficult to conform to at times.

Go fmt provides a good balance between formatting and style guide. For example, the following lines of code provide the same end result :

# Example 1:
	i, j, k := "Go", 42, true

# Example 2
	i := "go"
	j := 42
	k := true

Go.fmt won't try to selectively choose one style over the other however it will format the code respective to the style chosen e.g. 

# Format 1 
	i, j, k := "Go",       42, true

Will be save to this once go.fmt is ran :

# After g0.fmt
	i, j, k := "Go", 42, true

