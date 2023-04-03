<template>
  <v-app>
    <HeaderPrincipal/>
    <v-main>
      <RouterView></RouterView>
    </v-main>
  </v-app>
</template>

<script>
import HeaderPrincipal from './components/Header.vue'

export default {
  name: 'App',

  components: {
    HeaderPrincipal,
  },

  data: () => ({
    //
  }),
}
</script>
