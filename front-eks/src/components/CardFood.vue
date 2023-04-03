<template>
    <v-card
      class="mx-auto"
      max-width="344"
    >
      <v-img
        :src="image" alt="food image"
        height="200px"
        cover
      ></v-img>
  
      <v-card-title>
        {{ name_food }}
      </v-card-title>
  
      <v-card-subtitle>
        ${{ price }}
      </v-card-subtitle>
  
      <v-card-actions>
        <v-btn
          color="orange-lighten-2"
          variant="text"
        >
          Descripción
        </v-btn>
  
        <v-spacer></v-spacer>
  
        <v-btn
          :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'"
          @click="show = !show"
        ></v-btn>
      </v-card-actions>
  
      <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>
          <v-card-text>
            {{ category }}
          </v-card-text>
          <v-card-text>
            {{ description }}
          </v-card-text>
        </div>
      </v-expand-transition>
    </v-card>
  </template>
  
  <script>
  export default {
    name: 'CardFood',
    props: {
      name_food: {
        type: String,
        required: true,
      },
      price: {
        type: Number,
        required: true,
      },
      description: {
        type: String,
        required: true,
      },
      category: {
        type: String,
        required: true,
      },
      food_id: {
        type: Number,
        required: true,
      },
      image: {
        type: String,
        required: true,
      },
    },
    data: () => ({
      show: false,
    }),
  };
  </script>
  
  <style>
  .card {
    display: flex;
    flex-direction: column;
    background-color: #f5f5f5;
    border-radius: 10px;
    box-shadow: 0 3px 6px rgba(0, 0, 0, 0.3);
    overflow: hidden;
    width: 300px;
  }

  </style>
  