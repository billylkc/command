package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"d"},
	Short:   "[d] For development purposes.",
	Long:    `[d] For development purposes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := `
Logistic regression (https://en.wikipedia.org/wiki/logistic_regression)
In statistics, the logistic model (or logit model) is used to model the probability of a certain class or event existing such as pass/fail, win/lose, alive/dead or healthy/sick.  This can be extended to model several classes of events such as determining whether an image contains a cat, dog, lion, etc.  Each object being detected in the image would be assigned a probability between 0 and 1, with a sum of one.
logistic regression is a statistical model that in its basic form uses a logistic function to model a binary dependent variable, although many more complex extensions exist. In regression analysis, logistic regression[1] (or logit regression) is estimating the parameters of a logistic model (a form of binary regression). Mathematically, a binary logistic model has a dependent variable with two possible values, such as pass/fail which is represented by an indicator variable, where the two values are labeled "0" and "1". In the logistic model, the log-odds (the logarithm of the odds) for the value labeled "1" is a linear combination of one or more independent variables ("predictors"); the independent variables can each be a binary variable (two classes, coded by an indicator variable) or a continuous variable (any real value). The corresponding probability of the value labeled "1" can vary between 0 (certainly the value "0") and 1 (certainly the value "1"), hence the labeling; the function that converts log-odds to probability is the logistic function, hence the name. The unit of measurement for the log-odds scale is called a logit, from logistic unit, hence the alternative names. Analogous models with a different sigmoid function instead of the logistic function can also be used, such as the probit model; the defining characteristic of the logistic model is that increasing one of the independent variables multiplicatively scales the odds of the given outcome at a constant rate, with each independent variable having its own parameter; for a binary dependent variable this generalizes the odds ratio.
In a binary logistic regression model, the dependent variable has two levels (categorical). Outputs with more than two values are modeled by multinomial logistic regression and, if the multiple categories are ordered, by ordinal logistic regression (for example the proportional odds ordinal logistic model[2]). The logistic regression model itself simply models probability of output in terms of input and does not perform statistical classification (it is not a classifier), though it can be used to make a classifier, for instance by choosing a cutoff value and classifying inputs with probability greater than the cutoff as one class, below the cutoff as the other; this is a common way to make a binary classifier. The coefficients are generally not computed by a closed-form expression, unlike linear least squares; see § Model fitting. The logistic regression as a general statistical model was originally developed and popularized primarily by Joseph berkson,[3] beginning in Berkson (1944) harvtxt error: no target: CITEREFBerkson1944 (help), where he coined "logit"; see § History.
`
		fmt.Println(s)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
