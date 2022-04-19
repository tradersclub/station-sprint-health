type Person = {
  personName: string;
  birthday: number;
};

const person = {
  personName: "Juliana",
  birthday: 956026800000,
};

const data = [
  {
    personName: "ana",
    birthday: 956026800000,
  },
  {
    personName: "Juliana",
    birthday: 956026800000,
  },
];

function isPalindrome(person: string) {
  if (!person) {
    console.log("needs a person");
  }

  const reversedName = person.split("").reverse().join("");

  let isPalindrome: boolean =
    reversedName.toLowerCase() === person.toLowerCase();

  return isPalindrome;
}

function isHigherThanEightteen(birthday: number) {
  const birthYear = new Date().getFullYear() - new Date(birthday).getFullYear();
  return birthYear >= 18;
}

function personsCheck(persons: Person[]) {
  persons.filter((person: Person) => {
    if (
      isPalindrome(person.personName) &&
      isHigherThanEightteen(person.birthday)
    ) {
      return `${person.personName} is isHigherThanEightteen`;
    }
    return "not isHigherThanEightteen";
  });
}

console.log(personsCheck(data));
