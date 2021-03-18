import Vue from 'vue'
import VueResource from 'vue-resource'

import App from './App.vue'

import 'bootstrap/dist/js/bootstrap.js'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'


Vue.use(VueResource)


new Vue({
	render: h => h(App),
}).$mount('#app')
