<template>
	<div>
		<ModelAbilities
			v-for="value in models"
			v-bind:model="value"
			:key="value.id"
			:abilitiesList="abilitiesList"
			v-on:update="getAbilities"
		></ModelAbilities>
	</div>
</template>

<script>
import ModelAbilities from "./model_abilities.vue"
export default {
	name: "Abilities",
	props: ["ref_id"],
	components: { ModelAbilities },
	watch: {
		ref_id: function(newVal) {
			this.getModels(newVal)
		},
	},
	created: function() {
		this.getAbilities()
		this.getModels(this.ref_id)
	},
	data() {
		return {
			abilitiesList: [],
			models: [],
		}
	},
	methods: {
		getAbilities: function() {
			this.$http.get(process.env.VUE_APP_API_ENDPOINT + `/abilities?lang=${this.$language}`).then(function(res) {
				console.debug(res)
				this.abilitiesList = res.data
			})
		},
		getModels: function(refID) {
			this.models = []
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + "/ref/" + refID + "/model?lang=" + this.$language)
				.then(function(res) {
					console.debug(res)
					this.models = res.data
				})
		},
	},
}
</script>

<style scoped>
</style>
