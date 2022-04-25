#include "../utils.h"
#include <stdbool.h>

typedef struct {
  char *firstName;
  char *lastName;
} fullName;

typedef struct {
  char *Month;
  char *Day;
  char *Year;
} Date;

typedef struct {
  Date lastPaid;
} Subscription;

typedef struct {
  fullName name;
  char *signUpDate;
  bool paying;
  bool volunteer;
  Subscription subscription;
} user;

inline user newUser() {
  char *firstName = askForInput("What is your first name?");
  char *lastName = askForInput("What is your last name?");

  fullName fullname;
  fullname.firstName = firstName;
  fullname.lastName = lastName;

  bool paying =
      stringToBool(askForInput("Would you like to pay now? (yes/no)"));
  bool volunteer =
      stringToBool(askForInput("Would you like to volunteer? (yes/no)"));
};

inline void saverUser(user user){};

inline void getAllUsers(user *user) {
  for (int i = 0; i < sizeof(*user); i++) {
  }
};

inline user checkUserDup(user user) {}
