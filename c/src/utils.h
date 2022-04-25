#include <stdbool.h>
#include <stdio.h>
#include <string.h>

inline char *askForInput(char *message) {
  char *answer;

  printf("%s", message);
  scanf("%s", answer);

  return answer;
};

inline bool stringToBool(char *string) {
  if (strcmp("yes", string) || strcmp("y", string)) {
    return true;
  } else if (strcmp("no", string) || strcmp("n", string)) {
    return false;
  }

  return false;
}
