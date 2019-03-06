<template>
	<div class="w-100  mt-4">
		<CardAbilities :abilitiesList="abilitiesList" :id="card.id" v-on:new="newAbility" v-on:update="updateAbility"></CardAbilities>
		<ModelAbilities v-for="value in card.models" v-bind:model="value" :key="value.id" :abilitiesList="abilitiesList" v-on:new="newAbility" v-on:update="updateAbility"></ModelAbilities>
	</div>
</template>

<script>
import CardAbilities from "./card_abilities.vue";
import ModelAbilities from "./model_abilities.vue";
export default {
	name: "Abilities",
	props: ["card"],
	components: {
		CardAbilities,
		ModelAbilities
	},
	watch: {},
	created: function() {
		this.get();
	},
	data() {
		return {
			abilitiesList: []
		};
	},
	methods: {
		get: function() {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/abilities")
				.then(function(res) {
					console.log(res);
					this.abilitiesList = res.data;
				});
		},
		newAbility: function(ability) {
			this.abilitiesList.push(ability);
		},
		updateAbility: function() {
			this.get();
		},
	}
};
</script>

<style scoped>
</style>
