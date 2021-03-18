import Vue from 'vue'
import VueResource from 'vue-resource'
import Autocomplete from 'v-autocomplete'
import CountryFlag from 'vue-country-flag'
import TextComplete from 'v-textcomplete'

import App from './App.vue'

import 'bootstrap/dist/js/bootstrap.js'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'
import './custom.scss'

export const EventBus = new Vue()

Vue.use(VueResource)
Vue.use(Autocomplete)
Vue.use(CountryFlag)
Vue.component('v-textcomplete', TextComplete)

let globalData = new Vue({
	data: { $language: 'FR' },
})
Vue.mixin({
	computed: {
		$language: {
			get: function() {
				return globalData.$data.$language
			},
			set: function(newVal) {
				globalData.$data.$language = newVal
			},
		},
	},
})

new Vue({
	render: h => h(App),
}).$mount('#app')
