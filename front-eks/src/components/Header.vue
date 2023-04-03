<template>
    <v-toolbar app class="warning">
    <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title class="clean-text"><router-link to="/">EKS-Restaurant</router-link></v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn class="succes clean-text">
        <router-link to="/Indian" >Indian</router-link>
    </v-btn>
    <v-btn class="succes clean-text">
        <router-link to="/Mexican">Mexican</router-link>
    </v-btn>
    <v-btn class="succes clean-text">
        <router-link to="/Chinese">Chinese</router-link>
    </v-btn>
</v-toolbar>

<v-navigation-drawer app v-model="drawer" temporary dark> 
    <v-layout mt-4 aling-center>
      <div class="d-flex flex-column menu-lateral">
        <v-flex>
          <v-avatar>
            <img clasname="h-90 w-90" src="https://img.icons8.com/ios-filled/100/null/name.png" alt="">
          </v-avatar>
        </v-flex>
        <v-flex>
          <p clase="white--text mt-3 headline">Usuario Prueba</p>
        </v-flex>
      </div>		
	</v-layout>
  </v-navigation-drawer>
</template>

<script>

export default {
  name: 'HeaderPrincipal',

  data: () => {

    return{
      drawer: false,
    };
  },

}
</script>
<style>
.clean-text a:-webkit-any-link{
    text-decoration: none;
    color: #050505;
    font-size: 20px;
}

.menu-lateral{
    
}
</style>