<template>
	<div id="app" class="container">
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
				<button type="submit" class="btn btn-primary" @click="setCard(card)">Go</button>
				<button type="submit" class="btn btn-primary" @click="setCard(0)">New Card</button>
			</form>
		</div>
		<div class="row">
			<Content v-if="cardID !== null" :id="cardID" :factions="factions" :categories="categories"/>
		</div>
	</div>
</template>

<script>
import Content from "./components/content.vue"
export default {
	name: "app",
	components: {
		Content
	},
	data() {
		return {
			factions: [],
			faction: null,
			categories: [],
			category: null,
			cards: [],
			card: null,
			cardID: null
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
		setCard: function(card) {
			this.cardID = card;
		}
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
	}
};
</script>

<style>
#app {
	font-family: "Avenir", Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
}

body,
html {
	height: 100%;
}

.h-head {
	height: 10vh;
}
.h-content {
	height: 90vh;
}
</style>
