<template>
	<div class="w-100">
		<div v-if="id==='0'" class="alert alert-warning" role="alert">Creating a new model...</div>
		<form>
			<div class="form-group row">
				<label class="col-2 col-form-label">Faction</label>
				<div class="col-10">
					<select v-model="model.faction_id" class="form-control">
						<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
					</select>
				</div>
				<label class="col-2 col-form-label">Category</label>
				<div class="col-10">
					<select v-model="model.category_id" class="form-control">
						<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
					</select>
				</div>

				<label class="col-2 col-form-label">Fullname</label>
				<div class="col-10">
					<input
						v-model="model.name"
						type="text"
						class="form-control"
						placeholder="Fyanna, Torment of Everblight"
					>
				</div>
				<label class="col-2 col-form-label">Shortname</label>
				<div class="col-10">
					<input v-model="model.shortname" type="text" class="form-control" placeholder="Fyanna2">
				</div>
				<label class="col-2 col-form-label">Properties</label>
				<div class="col-10">
					<input
						v-model="model.properties"
						type="text"
						class="form-control"
						placeholder="Blighted strider warlock"
					>
				</div>
			</div>
		</form>

		<div class="row">
			<label class="col-2 col-form-label">Statline</label>
			<div class="col-10 text-left">
				<form class="form-inline statline">
					<input v-model="model.spd" type="text" class="form-control" placeholder="SPD">
					<input v-model="model.str" type="text" class="form-control" placeholder="STR">
					<input v-model="model.mat" type="text" class="form-control" placeholder="MAT">
					<input v-model="model.rat" type="text" class="form-control" placeholder="RAT">
					<input v-model="model.def" type="text" class="form-control" placeholder="DEF">
					<input v-model="model.arm" type="text" class="form-control" placeholder="ARM">
					<input v-model="model.cmd" type="text" class="form-control" placeholder="CMD">
					
					<input v-model="model.cost" type="text" class="form-control" placeholder="Cost">
					<input v-model="model.fa" type="text" class="form-control" placeholder="FA">
				</form>
			</div>
		</div>

		<div class="row">
			<label class="col-2 col-form-label">Advantages</label>
			<div class="col-10 text-left">
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.advance_deploy">
					<label class="form-check-label">Advance deploy</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.amphibious">
					<label class="form-check-label">Amphibious</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.arc_node">
					<label class="form-check-label">Acr node</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.assault">
					<label class="form-check-label">Assault</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.cavalry">
					<label class="form-check-label">Cavalry</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.cma">
					<label class="form-check-label">CMA</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.cra">
					<label class="form-check-label">CRA</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.construct">
					<label class="form-check-label">Construct</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.eyeless_sight">
					<label class="form-check-label">Eyeless sight</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.flight">
					<label class="form-check-label">Flight</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.gunfighter">
					<label class="form-check-label">Gunfighter</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.incorporeal">
					<label class="form-check-label">Incorporeal</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.immune_corrosion">
					<label class="form-check-label">Immune corrosion</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.immune_electricity">
					<label class="form-check-label">Immune electricity</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.immune_fire">
					<label class="form-check-label">Immune fire</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.immune_frost">
					<label class="form-check-label">Immune frost</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.jackmarshal">
					<label class="form-check-label">Jack Marshal</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.officer">
					<label class="form-check-label">Officer</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.parry">
					<label class="form-check-label">Parry</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.pathfinder">
					<label class="form-check-label">Pathfinder</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.soulless">
					<label class="form-check-label">Soulless</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.stealth">
					<label class="form-check-label">Stealth</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.tought">
					<label class="form-check-label">Tought</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="model.undead">
					<label class="form-check-label">Undead</label>
				</div>
			</div>
		</div>


		<div class="row">
			<label class="col-2 col-form-label"></label>
			<div class="col-10">
				<form class="form-inline statline">
					<v-autocomplete
						:items="items"
						v-model="weapon"
						:get-label="getLabel"
						@update-items="updateItems"
						:component-item="template"
						:auto-select-one-item="false"
						@item-selected="inputWeapon"
					></v-autocomplete>
					<input v-model="weapon.type" type="text" class="form-control" placeholder="Type">
					<input v-model="weapon.rng" type="text" class="form-control" placeholder="RNG">
					<input v-model="weapon.pow" type="text" class="form-control" placeholder="POW">
					<input v-model="weapon.rof" type="text" class="form-control" placeholder="ROF">
					<input v-model="weapon.aoe" type="text" class="form-control" placeholder="AOE">
					,
					res :{{weapon}}
					res :{{search}}
				</form>
			</div>
		</div>

		<div class="row" id="example-3">
			<input type="checkbox" id="jack" value="Jack" v-model="checkedNames">
			<label for="jack">Jack</label>
			<input type="checkbox" id="john" value="John" v-model="checkedNames">
			<label for="john">John</label>
			<input type="checkbox" id="mike" value="Mike" v-model="checkedNames">
			<label for="mike">Mike</label>
			<br>
			<span>Checked names: {{ checkedNames }}</span>
		</div>

		<div class="row">
			<label class="col-2 col-form-label"></label>
			<div class="col-10">
				<button type="submit" class="form-control btn btn-primary" @click="save()">Save</button>
			</div>
		</div>
	</div>
</template>

<script>
import ItemTemplate from "./ItemTemplate.vue";

export default {
	name: "Content",
	props: ["id", "factions", "categories"],
	watch: {
		id: function(newVal, oldVal) {
			console.log(newVal, oldVal);
			// watch it
			this.model = {};
			this.$http
				.get("http://localhost:9901/models/" + this.id)
				.then(function(res) {
					console.log(res);
					this.model = res.data;
				})
				.catch(function(err) {
					console.log(err);
				});
		}
	},
	data() {
		return {
			model: {},
			items: [],
			item: null,
			search: null,
			weapons: [
				{ id: 1, name: "Antiphon" },
				{ id: 2, name: "Rapture" },
				{ id: 3, name: "Rapace" }
			],
			weapon: {},
			template: ItemTemplate,
			checkedNames:["Mike"],
		};
	},
	methods: {
		save: function() {
			this.$http
				.put("http://localhost:9901/models/" + this.id, this.model)
				.then(function(res) {
					console.log(res);
					console.log(res.data);
					if (res.status === 201) {
						this.id = res.data;
					}
				});
		},
		getLabel: function(item) {
			if (!item) {
				return;
			}
			console.log("item", item);
			return item.name;
		},
		updateItems(text) {
			this.items = this.weapons.filter(item =>
				item.name.toLowerCase().startsWith(text.toLowerCase())
			);
			this.search = text;
		},
		inputWeapon(text) {
			console.log("input", text);
			this.search = text.name;
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.statline input {
	max-width: 4rem;
}
</style>
