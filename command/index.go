package command

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"time"
)

const (
	startDateFlagDefaultValue  = "19000101"
	endDateFlagDefaultValue    = "21001231"
	stepFlagDefaultValue       = 24 * time.Hour
	dateFormatFlagDefaultValue = "20060102"
)

func getDateFlagValue(flags *pflag.FlagSet, flagName string, dateFormat string) (dateValue time.Time, err error) {
	flagValue, err := flags.GetString(flagName)
	if err != nil {
		return time.Time{}, errors.Wrap(err, `command/generation->runGeneration(): GetString("`+flagName+`")`)
	}

	dateValue, err = time.Parse(dateFormat, flagValue)
	if err != nil {
		return time.Time{}, errors.Wrap(err, `command/generation->runGeneration(): time.Parse() for "`+flagName+`"`)
	}

	return
}

func runGeneration(cmd *cobra.Command, args []string) (err error) {
	dateFormatFlagValue, err := cmd.Flags().GetString("date-format")
	if err != nil {
		return errors.Wrap(err, `command/generation->runGeneration(): GetString("date-format")`)
	}

	startDate, err := getDateFlagValue(cmd.Flags(), "start-date", dateFormatFlagValue)
	if err != nil {
		return errors.Wrap(err, `command/generation->runGeneration(): getDateFlagValue() for "start-date")`)
	}

	endDate, err := getDateFlagValue(cmd.Flags(), "end-date", dateFormatFlagValue)
	if err != nil {
		return errors.Wrap(err, `command/generation->runGeneration(): getDateFlagValue() for "start-date")`)
	}

	rangeStepFlagValue, err := cmd.Flags().GetDuration("range-step")
	if err != nil {
		return errors.Wrap(err, `command/generation->runGeneration(): GetDuration("range-step")`)
	}

	currentDate := startDate
	for ; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.Add(rangeStepFlagValue) {
		fmt.Println(currentDate.Format(dateFormatFlagValue))
	}

	return errors.Wrap(err, "generation.runGeneration()")
}

func GetGenerationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "date-range-generator",
		Short: "Generates all dates in a range",
		Args:  cobra.ExactArgs(0),
		RunE:  runGeneration,
	}

	cmd.Flags().StringP("start-date", "s", startDateFlagDefaultValue, "The date to start")
	cmd.Flags().StringP("end-date", "e", endDateFlagDefaultValue, "The date to end")
	cmd.Flags().DurationP("range-step", "t", stepFlagDefaultValue, "The step to move in the date range")
	cmd.Flags().StringP("date-format", "f", dateFormatFlagDefaultValue, "Date format from this date: 2006-Jan-02 Monday 03:04:05 PM MST -07:00")

	return cmd
}
