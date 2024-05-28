document.getElementById('shorten-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const url = document.getElementById('url').value;
    const resultDiv = document.getElementById('result');

    // Clear previous results
    resultDiv.innerHTML = '';

    fetch('http://127.0.0.1:8080/links', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: url })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.short_url) {
                const shortenedUrl = 'http://localhost:8080/links/' + data.short_url;
                resultDiv.innerHTML = `<p>Shortened URL: <a href="${shortenedUrl}" target="_blank">${shortenedUrl}</a></p>`;
            } else {
                resultDiv.innerHTML = `<p style="color: red;">Error: ${data.error}</p>`;
            }
        })
        .catch(error => {
            resultDiv.innerHTML = `<p style="color: red;">Unable to shorten the URL. Please try again later.</p>`;
        });
});