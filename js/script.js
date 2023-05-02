// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  // This function creates a random number, either positive or negative
  let randomNumber = 0.00
  // input
  const positiveButtonChecked = document.getElementById("positive-number").checked

  // process
  if (positiveButtonChecked) {
    // positive
    randomNumber = Math.floor(Math.random() * 6)
  } else {
    // negative
    randomNumber = Math.floor(Math.random() * -6)
  }

  // output
  document.getElementById('random-number').innerHTML = randomNumber
}
