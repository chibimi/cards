<template>
	<div class="w-100">
		<h4>Abilities</h4>
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
				.get("http://localhost:9901/abilities")
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
