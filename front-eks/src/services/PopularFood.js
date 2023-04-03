import Api from "@/services/Api"

export default {
    getPopularFood() {
        return Api().get('/foods')
    }
}