const button = document.querySelector('.submit-button');

function init() {
    button.addEventListener("click", handleClick);
}

function handleClick(e) {
    console.log(`This event: ${e} happened.`);
    fetch('http://localhost:8080/api/data', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log(data.message); // do something with the data
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        })
}

init();