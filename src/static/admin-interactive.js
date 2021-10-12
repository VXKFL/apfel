document.addEventListener('alpine:init', () => {
    Alpine.store('data', {
        init() {
            this.users = ["User1", "User2", "User3"]
        },

        users: [],
    })
})
