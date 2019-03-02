<template>
	<div class="w-100">
		<div class="row my-4">
			<h2 class="text-left col-6">{{card.name}}</h2>
			<div class="col-4" >
				<div v-if="alert" class="alert py-2" v-bind:class="{ 'alert-success': alert_success,'alert-danger': !alert_success}">{{alert}}</div>
			</div>
			<div class="col-1 pl-0">
				<button v-if="update" type="submit" class="form-control btn btn-primary" @click="save(card)">Save</button>
				<button v-else type="submit" class="form-control btn btn-success" @click="update = !update">Update</button>
			</div>
			<div class="col-1 pl-0">
				<button type="submit" class="form-control btn btn-danger" @click="remove(card)">Delete</button>
			</div>
		</div>
		<div v-if="update">
			<div class="form-group row">
				<div class="col-6">
					<div class="row">
						<label class="col-form-label col-3">Name</label>
						<input v-model="card.name" type="text" class="form-control col-8" placeholder="Fyanna, Torment of Everblight">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Tag</label>
						<input v-model="card.properties" type="text" class="form-control col-8" placeholder="Unité de la légion">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Faction</label>
						<select v-model="card.faction_id" class="form-control col-8">
							<option v-for="f in factions" :key="f.id" v-bind:value="f.id">{{f.name}}</option>
						</select>
					</div>
					<div class="row">
						<label class="col-form-label col-3">Type</label>
						<select v-model="card.category_id" class="form-control col-8">
							<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
						</select>
					</div>
				</div>

				<div class="col-6">
					<div class="row">
						<label class="col-form-label col-3">FA</label>
						<input v-model="card.fa" type="text" class="form-control col-2">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Cost</label>
						<input v-model="card.cost" type="text" class="form-control col-2">
						<label v-if="card.category_id===5" class="col-form-label col-4">Cost max</label>
						<input v-if="card.category_id===5" v-model="card.cost_max" type="text" class="form-control col-2">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Nb model</label>
						<input v-model="card.models" type="text" class="form-control col-2">
						<label v-if="card.category_id===5" class="col-form-label col-4">Nb model max</label>
						<input v-if="card.category_id===5" v-model="card.models_max" type="text" class="form-control col-2" placeholder>
					</div>
					<div class="row">
						<label class="col-form-label col-3">Main ID</label>
						<input v-model="card.main_card_id" type="text" class="form-control col-2">
						<label class="col-form-label col-4">Status</label>
						<select v-model="card.status" class="form-control col-3">
							<option value="wip">WIP</option>
							<option value="tbv">A valider</option>
							<option value="done">Terminée</option>
						</select>
					</div>
				</div>
			</div>
		</div>

		<div v-else class="text-left row">
			<div class="col-4">
				<div>
					<b>Name:</b>
					{{card.name}}
				</div>
				<div>
					<b>Tag:</b>
					{{card.properties}}
				</div>
				<div>
					<b>Faction:</b>
					{{card.faction_id}}
				</div>
				<div>
					<b>Type:</b>
					{{card.category_id}}
				</div>
			</div>
			<div class="col-4">
				<div>
					<b>FA:</b>
					{{card.fa}}
				</div>
				<div>
					<b>Cost:</b>
					{{card.cost}} for {{card.models}} models
				</div>
				<div v-if="card.category_id===5">
					<b>Cost max:</b>
					{{card.cost_max}} for {{card.models_max}} models
				</div>
				<div>
					<b>Main card ID:</b>
					{{card.main_card_id}}
				</div>
				<div>
					<b>Status:</b>
					{{card.status}}
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import { Factions, Categories } from "./const.js";
export default {
	name: "Ref",
	props: ["id", "faction", "category"],
	components: {},
	watch: {
		id: function(newVal) {
			if (newVal > 0) {
				this.get(newVal);
			}else{
				this.reset();
			}
			
		}
	},
	created: function() {
		if (this.id > 0) {
			this.get(this.id);
		}
	},
	data() {
		return {
			factions: Factions,
			categories: Categories,
			card: {
				status: "wip",
				category_id: this.category,
				faction_id: this.faction,
			},
			update: true,
			alert: "",
			alert_success: false,
		};
	},
	methods: {
		get: function(cardID) {
			this.$http
				.get("http://localhost:9901/cards/" + cardID)
				.then(function(res) {
					console.log(res);
					this.card = res.data;
					if (this.card.id > 0 && this.card.status !== "wip") {
						this.update = false;
					}
				});
		},
		save: function(card) {
			if (card.id == null) {
				card.id = 0;
			}
			this.alert = ""
			this.$http
				.put("http://localhost:9901/cards/" + card.id, card)
				.then(function(res) {
					this.alert = "save success"
					this.alert_success = true
					if (res.status === 201) {
						this.$emit('new_card',res.data)
					}
					if (this.card.id > 0 && this.card.status !== "wip") {
						this.update = false;
					}
				}).catch(function(err){
					this.alert = "error: "+err.data
					this.alert_success = false
				});
		},
		remove: function(card) {
						this.alert = ""
			this.$http
				.delete("http://localhost:9901/cards/" + card.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
												this.$emit('remove_card', card.id)
					}
				}).catch(function(err){
					this.alert = "error: "+err
					this.alert_success = false
				});
		},
		reset: function(){
			this.update = true;
			this.card = {
				status: "wip",
				faction_id:this.faction,
				category_id:this.category
			};
			this.alert= ""
			this.alert_succes= false
		},
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
