#include <stdio.h>

int main(void)
{
	int MULTIPLIER = 12;
	int minutes;
	int bottles;

	printf("Enter how many minutes you shower: ");
	scanf("%d", &minutes);

	bottles = MULTIPLIER * minutes;

	printf("You entered: %d\n", bottles);
}

