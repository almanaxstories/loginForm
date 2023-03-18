console.log('JS-file connected successfully!');

window.addEventListener("load", ()=>{
  let loginButton = document.getElementById("loginButton");
  loginButton? console.log("Login button fetched successfully!") : console.log("Login button is not fetched!");
  let signinForm = document.getElementById("signinForm");
  signinForm? console.log("Signin form fetched successfully!") : console.log("Signin form is not fetched!");
  console.log(loginButton);
  console.log(signinForm);

  loginButton.addEventListener('click', event => {
      //event.preventDefault();
      console.log("Button works!");
      
      const formData = new FormData(signinForm);
      const objectFromInputs = Object.fromEntries(formData);
      
      const xhr = new XMLHttpRequest();
      xhr.open("POST", "127.0.0.1:8090");
      xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
      const body = JSON.stringify(objectFromInputs);
      /*xhr.onload = () => {
          if (xhr.readyState == 4 && xhr.status == 201) {
            console.log(JSON.parse(xhr.responseText));
          } else {
            console.log(`Error: ${xhr.status}`);
          }
        };*/
      xhr.send(body);
      console.log(objectFromInputs);
      console.log(body);
  });
})
