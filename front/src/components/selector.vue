<template>
	<div class="row">
		<form v-on:submit.prevent class="form-inline">
			<label>Faction</label>
			<select v-model="faction" class="form-control" @change="changeFaction">
				<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
			</select>
			<label>Category</label>
			<select v-model="category" class="form-control" @change="changeCategory">
				<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
			</select>
			<label>Card</label>
			<select v-model="card" class="form-control">
				<option v-for="c in cards" :key="c.id" :value="c.id">[{{c.status}}] #{{c.id}} {{c.name}}</option>
			</select>
			<button type="submit" class="btn btn-primary" @click="$emit('select_card', card)">Go</button>
			<button type="submit" class="btn btn-primary" @click="$emit('select_card', 0)">New Card</button>
		</form>
	</div>
</template>

<script>
import { Factions, Categories } from "./const.js";
export default {
	name: "Selector",
	components: {},
	data() {
		return {
			factions: Factions,
			faction: 11,
			categories: Categories,
			category: 5,
			cards: [],
			card: null
		};
	},
	methods: {
		getCards: function(faction, category) {
			if (!faction || !category) {
				return;
			}
			this.$http
				.get(
					process.env.VUE_APP_API_ENDPOINT+ "/cards?faction_id=" +
						faction +
						"&category_id=" +
						category
				)
				.then(function(res) {
					console.log(res);
					this.cards = res.data;
				})
				.catch(function(err) {
					console.log(err);
				});
		},
		changeFaction: function() {
			this.getCards(this.faction, this.category);
			this.$emit("change_faction", this.faction);
		},
		changeCategory: function() {
			this.getCards(this.faction, this.category);
			this.$emit("change_category", this.category);
		},
	},
	created: function() {
		this.getCards(this.faction, this.category);
				this.$emit("change_faction", this.faction);
		this.$emit("change_category", this.category);
	}
};
</script>

<style>
</style>
