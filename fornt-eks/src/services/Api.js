import axios from "axios";

export default() => {
    return axios.create({
        baseURL: "http://localhost:3000",
        headers: {
            'Accept': 'aplication/json',
            'Content-Type': 'aplication/json',
        }
    })
}