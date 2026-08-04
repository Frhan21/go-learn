/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export function cookingStatus(timerInMinutes) {
  if (timerInMinutes === undefined) {
    return 'You forgot to set the timer.'
  } 

  if (timerInMinutes === 0) {
    return "Lasagna is done."
  }

  return "Not done, please wait."
}

export function preparationTime(layers, numOfLayers = 2) {
  return layers.length * numOfLayers
}

export function quantities(layers) {
  let noodles = 0;
  let sauce = 0;

  for (let i = 0; i < layers.length; i++) {
    switch (layers[i]) {
      case 'noodles':
        noodles += 50;
        break;
      case 'sauce':
        sauce += 0.2;
        break;
    }
  }

  return {
    noodles,
    sauce,
  };
}

export function addSecretIngredient(friendList, myList) {
  const secretIngredient = friendList[friendList.length - 1];
  myList.push(secretIngredient);
}

export function scaleRecipe(recipe, numOfScaling) {

  var scaled = {}
  const portion = numOfScaling / 2
  
  for (const ingredient in recipe) {
    scaled[ingredient] = recipe[ingredient] * portion
  }

  return scaled
}
