import Vue from "vue";
import VueRouter from "vue-router";
import IndexPrincipal from "../components/Index.vue"
import IndianFood from "../components/IndianFood.vue"

Vue.use(VueRouter);

const routes = [
      {
        path: "/",
        name: "home",
        component: IndexPrincipal
      },
      {
        path: "/ndian",
        name: "idian",
        component: IndianFood
      }
    ];

const router = new VueRouter({
        mode: "history",
        routes
      });
      
export default router;