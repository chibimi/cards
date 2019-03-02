<template>
	<div class="w-100">
		<h4>Card abilities</h4>
		<Ability v-for="(value,index) in abilities" v-bind:ability="value" :key="value.id" v-on:remove="removeAbility(value,index)"></Ability>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_card_ability"
				aria-expanded="false"
				aria-controls="new_card_ability"
			>New Card Ability</h5>
			<div class="collapse card-body p-1" id="new_card_ability">
				<Ability :ability="ability" v-on:add="addAbility"></Ability>
			</div>
		</div>
		<hr>
		<ModelAbilities v-for="(value) in models" :abilitiesList="abilitiesList" :model="value" :key="value.id"></ModelAbilities>
	</div>
</template>

<script>
import Ability from "./ability.vue";
import ModelAbilities from "./model_abilities.vue";
export default {
	name: "CardAbilities",
	props: ["id"],
	components: {
		Ability, ModelAbilities
	},
	watch: {
		id: function(newVal) {
			this.get(newVal)
		}
	},
	created: function() {
		this.get(this.id)
	},
	data() {
		return {
			models: [],
			abilities: [],
			ability: {},
			abilitiesList:[],
		};
	},
	methods: {
		get: function(cardID) {
			this.$http
				.get("http://localhost:9901/cards/" + cardID + "/abilities")
				.then(function(res) {
					this.abilities = res.data;
				});
			this.$http
				.get("http://localhost:9901/cards/" + cardID + "/models")
				.then(function(res) {
					this.models = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(
					"http://localhost:9901/cards/" +
						this.id +
						"/abilities/" +
						ability.id
				)
				.then(function(res) {
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(
					"http://localhost:9901/cards/" +
						this.id +
						"/abilities/" +
						ability.id +
						"?magical=" +
						ability.magical
				)
				.then(function(res) {
					if (res.status === 200) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
