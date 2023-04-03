<template>
    <v-container>
      <v-row class="text-center">
        
  
        <v-col class="mb-4">
          <h1 class="display-2 font-weight-bold mb-3">
            Chinese Menu
          </h1>
        </v-col>      
      </v-row>
    </v-container>
  
      <div>
        <CardFood
          name_food="Chinese"
          price="450"
          description="A traditional Mexican dish made with chili peppers and chocolate"
          category="Mexican"
          food_id="1"
          image="https://tipsparatuviaje.com/wp-content/uploads/2018/09/mole-comida-mexicana.jpg"
        />
      </div>
  
  </template>
  
  <script>
  
  import CardFood from "@/components/CardFood.vue";
  export default {
    name: 'ChineseFood',
  
    data: () => {
  
    },
  
    components: {
      CardFood,
    }
  }
  </script>
  