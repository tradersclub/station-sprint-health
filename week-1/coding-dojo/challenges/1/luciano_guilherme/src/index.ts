type Person = {
  personName: string;
  birthday: number;
};

const person = {
  personName: "Juliana",
  birthday: 956026800000,
};

function isPalindrome(person: Person) {
  const { personName } = person;
  if (!personName) {
    console.log("needs a person");
  }

  const reversedName = personName.split("").reverse().join("");
  console.log(reversedName)

  let isPalindrome: boolean =
    reversedName.toLowerCase() === personName.toLowerCase();

  return isPalindrome;
}

function isHigherThanEightteen(birthday: number) {
  // birthYear = new Date(birthday).getFullYear();
  // const actualYear = new Date().getFullYear()
  // return;
}

console.log(isPalindrome(person));
