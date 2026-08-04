const NUMBERS = [
  'no',
  'one',
  'two',
  'three',
  'four',
  'five',
  'six',
  'seven',
  'eight',
  'nine',
  'ten',
];

const capitalize = (str) => str.charAt(0).toUpperCase() + str.slice(1);

const generateVerse = (bottles) => {
  const currentNumStr = capitalize(NUMBERS[bottles]);
  const currentWord = bottles === 1 ? 'bottle' : 'bottles';

  const nextBottles = bottles - 1;
  const nextNumStr = NUMBERS[nextBottles];
  const nextWord = nextBottles === 1 ? 'bottle' : 'bottles';

  return [
    `${currentNumStr} green ${currentWord} hanging on the wall,`,
    `${currentNumStr} green ${currentWord} hanging on the wall,`,
    'And if one green bottle should accidentally fall,',
    `There'll be ${nextNumStr} green ${nextWord} hanging on the wall.`,
  ];
};

export const recite = (initialBottlesCount, takeDownCount) => {
  const lyrics = [];

  for (let i = 0; i < takeDownCount; i++) {
    const currentBottles = initialBottlesCount - i;
    const verse = generateVerse(currentBottles);

    if (i > 0) {
      lyrics.push('');
    }

    lyrics.push(...verse);
  }

  return lyrics;
};