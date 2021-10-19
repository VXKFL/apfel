document.addEventListener('alpine:init', () => {
    Alpine.store('data', {
        users: [],
    });
    loadAttendance();
});

function loadAttendance() {
    axios.get("api/event/1")
        .then((response) => {
            let users = response.data.map((userinfo) => userinfo.Name)
            Alpine.store("data").users = users
        })
}

document.getElementById('scan-button').addEventListener('click', () => {
    console.log("click")
    window.location.replace("/admin/scan")
});
