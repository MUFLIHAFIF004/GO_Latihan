function showAlert() {
	alert("Halo! Ini dari JavaScript di Go + Fiber!");
}const images = [
  "/static/img/img1.png",
  "/static/img/img2.png",
  "/static/img/img3.png"
];
let currentIndex = 0;

function changeImage() {
  const img = document.getElementById("slideshow");
  currentIndex = (currentIndex + 1) % images.length;
  img.src = images[currentIndex];
}

setInterval(changeImage, 3000); // Ganti setiap 3 detik
