import Vue from 'vue'
import VueResource from 'vue-resource'
import Autocomplete from 'v-autocomplete'
import CountryFlag from 'vue-country-flag'

import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap/dist/js/bootstrap.js'

export const EventBus = new Vue();

Vue.use(VueResource)
Vue.use(Autocomplete)
Vue.use(CountryFlag)

Vue.config.productionTip = false
Vue.prototype.$language = 'FR'
Vue.prototype.$change_language = function (l) {
	Vue.prototype.$language = l;
}

new Vue({
	render: h => h(App),
}).$mount('#app')
