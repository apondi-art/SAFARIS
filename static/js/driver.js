document.getElementById('driver-form').addEventListener('submit', function (e) {
    e.preventDefault();

    // Collect form data
    const formData = new FormData(this);
    const data = Object.fromEntries(formData.entries());

    // Here you would typically send this data to your server
    console.log('Form submitted with data:', data);

    // For demo purposes, we'll just show an alert
    alert('Thank you for your application! We will contact you soon.');

    // Clear the form
    this.reset();
});
