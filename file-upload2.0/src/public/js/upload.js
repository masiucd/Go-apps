let fileUpload = document.querySelector('input[type="file"]');
let customFileUpload = document.getElementById("custom-file-upload");

fileUpload.addEventListener("change", (event) => {
  let fileName = event.target.files[0].name;
  customFileUpload.querySelector("span").textContent = fileName;
});
