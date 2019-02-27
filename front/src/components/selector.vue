<template>
		<div class="row">
			<form v-on:submit.prevent class="form-inline">
				<label >Faction</label>
				<select v-model="faction" class="form-control" @change="getCards(faction, category)">
					<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
				</select>
				<label >Category</label>
				<select v-model="category" class="form-control" @change="getCards(faction, category)">
					<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
				</select>
				<label >Card</label>
				<select v-model="card" class="form-control">
					<option v-for="c in cards" :key="c.id" :value="c.id">[{{c.status}}] {{c.name}}</option>
				</select>
				<button type="submit" class="btn btn-primary" @click="$emit('selected', card)">Go</button>
				<button type="submit" class="btn btn-primary" @click="$emit('selected', 0)">New Card</button>
			</form>
		</div>
</template>

<script>
export default {
	name: "Selector",
	components: {
    },
	data() {
		return {
			factions: [],
			faction: 11,
			categories: [],
			category: 5,
			cards: [],
			card: null,
		};
	},
	methods: {
		getCards: function(faction, category) {
			if (!faction || !category){
				return
			}
			this.$http
			.get("http://localhost:9901/cards?faction_id="+faction+"&category_id="+category)
			.then(function(res) {
				console.log(res);
				this.cards = res.data;
			})
			.catch(function(err) {
				console.log(err);
			});
		},
	},
	created: function() {
		this.$http
			.get("http://localhost:9901/factions")
			.then(function(res) {
				console.log(res);
				this.factions = res.data;
			})
			.catch(function(err) {
				console.log(err);
			});
		this.$http
			.get("http://localhost:9901/categories")
			.then(function(res) {
				console.log(res);
				this.categories = res.data;
			})
			.catch(function(err) {
				console.log(err);
            });
        this.getCards(this.faction, this.category);
	}
};
</script>

<style>
</style>
