import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min'
import 'mdbootstrap/css/mdb.min.css'
const validatorPassword = require('zxcvbn');
const validatorEmail = require('email-validator');

const firstName = document.querySelector('#firstname');
const lastName = document.querySelector('#lastname');
const email = document.querySelector('#email');
const password = document.querySelector('#password');
const button = document.querySelector('#btn-sign');
const inputs = document.querySelectorAll('.form-control');

function checkAndRemoveError(input) {
  if (input.nextElementSibling) {
    input.nextElementSibling.remove();
  }
}

function containsInvalid(input) {
  return input.classList.contains('is-invalid');
}

function errorInput(input) {
  checkAndRemoveError(input);
  input.classList.add('is-invalid');
  const inputError = document.createElement('div');
  inputError.classList.add('invalid-feedback');
  input.parentNode.appendChild(inputError);
  return inputError;
}

function validateInput(input) {
  checkAndRemoveError(input);
  input.classList.remove('is-invalid');
  input.classList.add('is-valid');
}

function isEmptyInput(input) {
  input.value.trim();
  if (!input.value) {
    return 'the form must be filled';
  }

  return '';
}

function limitName(input) {
  if (input.value.length > 255) {
    return 'the maximum number of characters is 255';
  }

  return '';
}

function limitPassword(input) {
  if (input.value.length < 8 || input.value.length > 64) {
    return 'password must be between 8 and 64 characters';
  }

  return '';
}

function checkEmailValue(input) {
  if (!validatorEmail.validate(input.value)) {
    return 'E-mail is invalid';
  }

  return '';
}

function isError() {
  return [...inputs].some(input => {
    return containsInvalid(input);
  });
}

function showErrorMassage(input, errorMessage) {
  const errorBar = errorInput(input);
  errorBar.textContent = errorMessage;
}

function commentScore(pass) {
  let text = validatorPassword(pass).score;
  if (pass === '') {
    return ''
  }
  switch (text) {
    case 0:
      text = 'risky password';
      break;
    case 1:
      text = 'weak password';
      break;
    case 2:
      text = 'medium password';
      break;
    case 3:
      text = 'safely password';
      break;
    case 4:
      text = 'very unguessable: strong protection.';
      break;
    default:
      break;
  }

  return text;
}

function passwordCheckMassage() {
  const passwordValue = password.value.trim();
  if (password.nextElementSibling) {
    password.nextElementSibling.remove();
  }
  const passwordCheck = document.createElement('div');
  passwordCheck.textContent = commentScore(passwordValue);
  password.parentNode.appendChild(passwordCheck);
}

function handleSignUp() {
  const form = [
    {
      input: firstName,
      validators: [isEmptyInput, limitName],
    },
    {
      input: lastName,
      validators: [isEmptyInput, limitName],
    },
    {
      input: email,
      validators: [isEmptyInput, checkEmailValue],
    },
    {
      input: password,
      validators: [isEmptyInput, limitPassword],
    },
  ];

  form.forEach(element => {
    let errorMsg = element.validators.reduce((accumulator, validator) => {
      if (accumulator !== '') {

        return accumulator;
      }

      return validator(element.input);
    }, '');
    if (errorMsg !== '') {
      showErrorMassage(element.input, errorMsg);
    } else {
      validateInput(element.input);
    }
  });

  if (!isError()) {
    let formData = getFormData(form);
    sendRequest('POST', url, formData)
      .then(data => console.log(data))
      .catch(err => console.log(err));
    // return (window.location.href = 'result.html');
  }
}

const url = 'http://localhost:8080/form';

function sendRequest(method, url, body = null) {
  return fetch(url, {
    method: method,
    body: JSON.stringify(body),
    headers: {
      'Content-Type': 'application/json'
    },
  }).then(response => {
    return response.text()
  })
}

function getFormData(form) {
  let formData = [];
  form.forEach(element => {
    formData.push([element.input.id, element.input.value]);
  });

  return Object.fromEntries(formData);
}

function checkBusyEmail() {
  const emailValue = email.value.trim();
  let objEmail = {
    email: emailValue
  };
  if (email.nextElementSibling) {
    email.nextElementSibling.remove();
  }
  const emailMassage = document.createElement('div');
  emailMassage.id = "email-massage"
  getMassageForBusyEmail(objEmail)
  email.parentNode.appendChild(emailMassage);
}

function getMassageForBusyEmail(email) {
  sendRequest('POST', 'http://localhost:8080/possession', email)
    .then(data => {
      document.querySelector('#email-massage').innerHTML = data;
    })
}

const emailIn = document.querySelector('#emailIn');
const passwordIn = document.querySelector('#passwordIn');
const buttonIn = document.querySelector('#btn-signIn');


function handleSignIn() {
  const SignInform = [
    {
      input: emailIn,
      validators: [isEmptyInput, checkEmailValue],
    },
    {
      input: passwordIn,
      validators: [isEmptyInput, limitPassword],
    },
  ];

  SignInform.forEach(element => {
    let errorMsg = element.validators.reduce((accumulator, validator) => {
      if (accumulator !== '') {

        return accumulator;
      }

      return validator(element.input);
    }, '');
    if (errorMsg !== '') {
      showErrorMassage(element.input, errorMsg);
    } else {
      validateInput(element.input);
    }
  });

  if (!isError()) {
    let formData = getFormData(SignInform);
    sendRequest('POST', 'http://localhost:8080/login', formData)
      .then(data => console.log(data))
      .catch(err => console.log(err));
  }
}


if (button) {
  button.addEventListener('click', handleSignUp);
  password.addEventListener('keyup', passwordCheckMassage);
  email.addEventListener('keyup', checkBusyEmail);
  buttonIn.addEventListener('click', handleSignIn);
}

