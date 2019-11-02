<template>
	<div class="w-100">
		<h4 class="text-left">Card abilities</h4>
		<Ability
			v-for="(value,index) in abilities"
			v-bind:ability="value"
			:abilitiesList="abilitiesList"
			:key="value.id"
			v-on:remove="removeAbility(value,index)"
			v-on:update="updateAbility"
		></Ability>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_card_ability"
				aria-expanded="false"
				aria-controls="new_card_ability"
			>New Card Ability</h5>
			<div class="collapse card-body p-1" id="new_card_ability">
				<Ability :ability="ability" :abilitiesList="abilitiesList" v-on:add="addAbility" v-on:new="newAbility" v-on:update="updateAbility"></Ability>
			</div>
		</div>
		<hr>
	</div>
</template>

<script>
import Ability from "./ability.vue";
export default {
	name: "CardAbilities",
	props: ["ref_id", "abilitiesList"],
	components: {
		Ability
	},
	watch: {
		id: function(newVal) {
			this.get(newVal);
		},
		abilitiesList: function() {
			this.get(this.ref_id);
		}
	},
	created: function() {
		this.get(this.ref_id);
	},
	data() {
		return {
			abilities: [],
			ability: {}
		};
	},
	methods: {
		get: function(refID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/ability?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.abilities = res.data;
				});
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + this.ref_id + "/ability/" + ability.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + this.ref_id + "/ability/" + ability.id + "?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					if (res.status === 200) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
		},
		updateAbility: function() {
			this.$emit("update");
		},
		newAbility: function(ability) {
			this.$emit("new", ability);
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
