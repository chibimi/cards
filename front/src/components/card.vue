<template>
	<div class="w-100">
		<div v-if="card.id===0" class="alert alert-info" role="alert">Creating a new card...</div>
		<form v-on:submit.prevent>
			<div class="form-group row">
				<label class="col-2 col-form-label">Faction</label>
				<div class="col-10">
					<select v-model="card.faction_id" class="form-control">
						<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
					</select>
				</div>

				<label class="col-2 col-form-label">Category</label>
				<div class="col-10">
					<select v-model="card.category_id" class="form-control">
						<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
					</select>
				</div>

				<label class="col-2 col-form-label">Card Name</label>
				<div class="col-10">
					<input v-model="card.name" type="text" class="form-control" placeholder="Fyanna, Torment of Everblight">
				</div>

				<label class="col-2 col-form-label">Properties</label>
				<div class="col-10">
					<input v-model="card.properties" type="text" class="form-control" placeholder="Unité de la légion">
				</div>

				<label v-if="card.category_id !== 5" class="col-2 col-form-label">Cost</label>
				<div v-if="card.category_id !== 5" class="col-10">
					<input v-model="card.cost" type="text" class="form-control" placeholder="+28, 8-10-9, 8*1...">
				</div>

				<label v-if="card.category_id === 5" class="col-2 col-form-label">Cost Min</label>
				<div v-if="card.category_id === 5" class="col-4">
					<input v-model="card.cost" type="text" class="form-control" placeholder>
				</div>
				<label v-if="card.category_id === 5" class="col-2 col-form-label">Nb Min</label>
				<div v-if="card.category_id === 5" class="col-4">
					<input v-model="card.models" type="text" class="form-control" placeholder>
				</div>

				<label v-if="card.category_id === 5" class="col-2 col-form-label">Cost Max</label>
				<div v-if="card.category_id === 5" class="col-4">
					<input v-model="card.cost_max" type="text" class="form-control" placeholder>
				</div>
				<label v-if="card.category_id === 5" class="col-2 col-form-label">Nb Max</label>
				<div v-if="card.category_id === 5" class="col-4">
					<input v-model="card.models_max" type="text" class="form-control" placeholder>
				</div>

				<label class="col-2 col-form-label">FA</label>
				<div class="col-10">
					<input v-model="card.fa" type="text" class="form-control" placeholder="2, C, U">
				</div>

				<label class="col-2 col-form-label">Resource</label>
				<div class="col-4">
					<input v-model="card.resource" type="text" class="form-control" placeholder="Focus or Fury">
				</div>
				<label class="col-2 col-form-label">Threshold</label>
				<div class="col-4">
					<input v-model="card.threshold" type="text" class="form-control" placeholder="Threshold">
				</div>

				<label class="col-2 col-form-label">Life</label>
				<div class="col-10">
					<input v-model="card.damage" type="text" class="form-control" placeholder="Voir la doc pour le format">
				</div>

				<label class="col-2 col-form-label">Main Card ID</label>
				<div class="col-10">
					<input v-model="card.main_card_id" type="text" class="form-control" placeholder="pour les modèles avec plus d'une carte. Ex Rorsh & Brine">
				</div>

				<label class="col-2 col-form-label">Status</label>
				<div class="col-10">
					<select v-model="card.status" class="form-control">
						<option value="TODO">TODO</option>
						<option value="WIP">WIP</option>
						<option value="TBV">A valider</option>
						<option value="Done">Terminée</option>
					</select>
				</div>

				<label class="col-2 col-form-label"></label>
				<div class="col-8 mt-2">
					<button type="submit" class="form-control btn btn-primary" @click="save(card)">Save Card</button>
				</div>
				<div class="col-2 mt-2">
					<button type="submit" class="form-control btn btn-danger" @click="remove(card)">Delete Card</button>
				</div>
			</div>
		</form>
		<hr>
		<h3>Card abilities</h3>
		<div class="row">

			<label class="col-1 col-form-label"></label>
			<div class="col-11">
				<Ability v-for="(value,index) in abilities" :abilitiesList="abilitiesList" v-bind:ability="value" :key="value.id" v-on:remove="removeAbility(value,index)"></Ability>
				<Ability :ability="ability" :abilitiesList="abilitiesList" v-on:add="addAbility"></Ability>
				<!-- <Model v-bind:model="model" v-on:add="addModel()"></Model> -->
			</div>
		</div>

		<!-- <div class="row">
			<label class="col-1 col-form-label"></label>
			<div class="col-11">
				<Model v-for="(m,i) in models" v-bind:model="m" :key="m.id" v-on:remove="models.splice(i,1)"></Model>
				<Model v-bind:model="model" v-on:add="addModel()"></Model>
			</div>
		</div> -->
	</div>
</template>

<script>
import { Factions, Categories } from "./const.js";
import Ability from "./ability.vue";
import Model from "./model.vue";
export default {
	name: "Card",
	props: ["id", "abilitiesList"],
	components: {
		Model, Ability
	},
	watch: {
		id: function(newVal) {
			this.get(newVal);
		}
	},
	created: function() {
		this.get(this.id);
	},
	data() {
		return {
			factions: Factions,
			categories: Categories,
			card: {},
			abilities: [],
			ability: {},
			models: [],
			model: {
				advantages: []
			}
		};
	},
	methods: {
		save: function(card) {
			if (card.id == null) {
				card.id = 0;
			}
			this.$http
				.put("http://localhost:9901/cards/" + card.id, card)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						card.id = res.data;
					}
				});
		},
		remove: function(card) {
			this.$http
				.delete("http://localhost:9901/cards/" + card.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						card = {};
					}
				});
		},
		get: function(cardID) {
			this.card = {};
			if (cardID < 0) {
				cardID = 0;
			}
			this.$http
				.get("http://localhost:9901/cards/" + cardID)
				.then(function(res) {
					console.log(res);
					this.card = res.data;
				});
			// this.$http
			// 	.get("http://localhost:9901/models?card_id=" + cardID)
			// 	.then(function(res) {
			// 		console.log(res);
			// 		this.models = res.data;
			// 	});
			this.$http
				.get("http://localhost:9901/cards/" + cardID + "/abilities")
				.then(function(res) {
					console.log(res);
					this.abilities = res.data;
				});
		},
		removeAbility: function(ability, index) {
			console.log("remove abi", ability, index)
			this.$http
				.delete("http://localhost:9901/cards/" + this.card.id + "/abilities/" + ability.id)
				.then(function(res) {
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
			
		},
		addAbility: function(ability) {
			this.$http
				.put("http://localhost:9901/cards/" + this.card.id + "/abilities/" + ability.id)
				.then(function(res) {
					if (res.status === 200) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
			
		},
		// removeModel: function(index) {
		// 	this.models.splice(index, 1);
		// },
		// addModel: function() {
		// 	this.models.push(this.model);
		// 	this.model = { card_id: this.card.id, advantages: [] };
		// }
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
