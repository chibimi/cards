<template>
	<div class="w-100  mt-4">
		<CardAbilities :abilitiesList="abilitiesList" :ref_id="ref_id" v-on:new="newAbility" v-on:update="updateAbility"></CardAbilities>
		<ModelAbilities v-for="value in models" v-bind:model="value" :key="value.id" :abilitiesList="abilitiesList" v-on:new="newAbility" v-on:update="updateAbility"></ModelAbilities>
	</div>
</template>

<script>
import CardAbilities from "./card_abilities.vue";
import ModelAbilities from "./model_abilities.vue";
export default {
	name: "Abilities",
	props: ["ref_id"],
	components: {
		CardAbilities,
		ModelAbilities
	},
	watch: {
		ref_id: function(newVal) {
			this.getModels(newVal);
		}
	},
	created: function() {
		this.getAbilities();
		this.getModels(this.ref_id);
	},
	data() {
		return {
			abilitiesList: [],
			models: []
		};
	},
	methods: {
		getAbilities: function() {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + "/abilities?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.abilitiesList = res.data;
				});
		},
		updateAbility: function() {
			this.getAbilities();
		},
		newAbility: function(ability) {
			this.abilitiesList.push(ability);
		},
		getModels: function(refID) {
			this.models = [];
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/model?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.models = res.data;
				});
		},
	}
};
</script>

<style scoped>
</style>
