function fetchData() {
    fetch("/api/v1/resources")
        .then(response => response.json())
        .then(data => {
            document.getElementById("cpu-progress").value = data.cpu_usage_percentage
            document.getElementById("cpu-progress-description").innerText = `${data.cpu_usage_percentage}% of 100% used`
            document.getElementById("ram-progress").value = data.ram_usage_percent
            document.getElementById("ram-progress-description").innerText = `${data.ram_usage} GB of ${data.ram_total} GB used (${data.ram_usage_percent}%)`
            document.getElementById("storage-progress").value = data.storage_usage_percent
            document.getElementById("storage-progress-description").innerText = `${data.storage_usage} GB of ${data.storage_total} GB used (${data.storage_usage_percent}%)`
            document.getElementById("temp-progress").value = data.cpu_temperature
            document.getElementById("temp-progress-description").innerText = `${data.cpu_temperature}\u00B0C of 85\u00B0C (${data.cpu_temperature_percent}%)`
        })
        .catch(error => {
            console.error('Failed to fetch data:', error)
        });
}

window.onload = setInterval(fetchData, 1500);