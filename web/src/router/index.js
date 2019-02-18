import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

import index from '../pages/index'

export default new VueRouter({
    routes: [{
        path: '/',
        name: 'index',
        component: index
    }]
})