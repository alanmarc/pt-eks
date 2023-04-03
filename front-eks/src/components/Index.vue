<template>
  <v-container>
    <v-row class="text-center">
      

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Pupular Menu
        </h1>
      </v-col>      
    </v-row>
  </v-container>
    <div>
        <CardFood v-for="res in response" :key="res.key"
        :name_food="res.name_food"
        :price="res.price"
        :description="res.description"
        :category="res.category"
        :food_id="res.food_id"
        :image="res.image"
        />
    </div>

</template>

<script>
//import PopularFood from "@/services/PopularFood.js";
import CardFood from "@/components/CardFood.vue";
import axios from 'axios';
export default {
  name: 'IndexPrincipal',
  data: () => {
    return {
      response: []
    }
  },

  /*methods: {
    async getFoods () {
      const res = await PopularFood.getPopularFood()
      this.response = res.data
      console.log(this.response)
    }
  },*/

  mounted() {
    axios.get('http://localhost:3000/foods')
      .then(({data}) => {
        console.log(data)
        this.response = data
      })
      .catch(error => {
        console.log(error)
        // Aquí puedes manejar el error, por ejemplo mostrando un mensaje de error
      })
  },
  components: {
    CardFood,
  }
}
</script>
