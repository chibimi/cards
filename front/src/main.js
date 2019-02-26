import Vue from 'vue'
import VueResource from 'vue-resource'
import Autocomplete from 'v-autocomplete'

import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.css'

Vue.use(VueResource)
Vue.use(Autocomplete)

Vue.config.productionTip = false

new Vue({
	render: h => h(App),
}).$mount('#app')
