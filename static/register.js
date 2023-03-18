console.log('JS-file connected successfully!');

window.addEventListener("load", () =>{
    let signUpForm = document.getElementById("signupform");
    signUpForm? console.log("SignUpFrom was fetched successfully!") : console.log("SignUpForm was not fetched!");
    let signUpButton = document.getElementById("signupbtn");
    signUpButton? console.log("SignUp button fetched successfully!") : console.log("SignUp button was not fetched!");

    signUpButton.addEventListener("click", event =>{
        //event.preventDefault();
        const formData = new FormData(signUpForm);
        const objectFromInputs = Object.fromEntries(formData);
        const xhr = new XMLHttpRequest();
        xhr.open("POST", "127.0.0.1:8090");
        xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
        const body = JSON.stringify(objectFromInputs);
        xhr.send(body);
        console.log(objectFromInputs);
        console.log(body);
    })
})