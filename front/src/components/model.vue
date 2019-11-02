<template>
	<div class="w-100">
		<div class="row pt-2 pl-2">
			<h4
				class="text-left col-6 pt-2"
				data-toggle="collapse"
				v-bind:data-target="'#test_model_' + model.id"
				v-bind:aria-expanded="!model.id"
				v-bind:aria-controls="'test_model_' + model.id"
			>{{model.name}}</h4>
			<div class="col-4">
				<div v-if="alert" class="alert alert-error py-2" v-bind:class="{ 'alert-success': alert_success }">{{alert}}</div>
			</div>

			<div class="col-1 pl-0">
				<button v-if="model.id" type="submit" class="form-control btn btn-success" @click="save(model)">Update</button>
			</div>
			<div class="col-1 pl-0">
				<button v-if="!model.id" type="submit" class="form-control btn btn-primary" @click="save(model)">Save</button>
				<button v-if="model.id" type="submit" class="form-control btn btn-danger" @click="remove(model)">Delete</button>
			</div>
		</div>

		<div class="collapse" v-bind:id="'test_model_' + model.id" v-bind:class="{'show': !model.id }">
			<div class="row px-3">
				<div class="col-3 pt-4 pr-4">
					<div class="row">
						<label class="col-form-label col-4 px-0">Name</label>
						<input v-model="model.name" type="text" class="form-control col-8">
					</div>
					<div class="row">
						<label class="col-form-label col-4 px-0">Damage</label>
						<input v-model="model.damage" type="text" class="form-control col-8">
					</div>
					<div class="row">
						<label class="col-form-label col-4 px-0">Fury/Focus</label>
						<input v-model="model.resource" type="text" class="form-control col-8">
					</div>
					<div class="row">
						<label class="col-form-label col-4 px-0">Threshold</label>
						<input v-model="model.threshold" type="text" class="form-control col-8">
					</div>
				</div>
				<div class="col-9 pl-3">
					<div class="form-group row my-0">
						<span class="col text-left">spd</span>
						<span class="col text-left">str</span>
						<span class="col text-left">mat</span>
						<span class="col text-left">rat</span>
						<span class="col text-left">def</span>
						<span class="col text-left">arm</span>
						<span class="col text-left">cmd</span>
						<span class="col text-left">base</span>
					</div>
					<div class="form-group row mt-0">
						<input v-model="model.spd" type="text" class="form-control col" placeholder="spd">
						<input v-model="model.str" type="text" class="form-control col" placeholder="str">
						<input v-model="model.mat" type="text" class="form-control col" placeholder="mat">
						<input v-model="model.rat" type="text" class="form-control col" placeholder="rat">
						<input v-model="model.def" type="text" class="form-control col" placeholder="def">
						<input v-model="model.arm" type="text" class="form-control col" placeholder="arm">
						<input v-model="model.cmd" type="text" class="form-control col" placeholder="cmd">
						<input v-model="model.base_size" type="text" class="form-control col" placeholder="case">
					</div>
					<div class="row">
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="advance_deploy">
							<label class="form-check-label">Advance deploy</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="amphibious">
							<label class="form-check-label">Amphibious</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="arc_node">
							<label class="form-check-label">Acr node</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="assault">
							<label class="form-check-label">Assault</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="cavalry">
							<label class="form-check-label">Cavalry</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="cma">
							<label class="form-check-label">CMA</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="cra">
							<label class="form-check-label">CRA</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="construct">
							<label class="form-check-label">Construct</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="eyeless_sight">
							<label class="form-check-label">Eyeless sight</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="flight">
							<label class="form-check-label">Flight</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="gunfighter">
							<label class="form-check-label">Gunfighter</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="incorporeal">
							<label class="form-check-label">Incorporeal</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="immunity_corrosion">
							<label class="form-check-label">Immune corrosion</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="immunity_electricity">
							<label class="form-check-label">Immune electricity</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="immunity_fire">
							<label class="form-check-label">Immune fire</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="immunity_frost">
							<label class="form-check-label">Immune frost</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="jackmarshal">
							<label class="form-check-label">Jack Marshal</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="officer">
							<label class="form-check-label">Officer</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="parry">
							<label class="form-check-label">Parry</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="pathfinder">
							<label class="form-check-label">Pathfinder</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="soulless">
							<label class="form-check-label">Soulless</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="stealth">
							<label class="form-check-label">Stealth</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="tought">
							<label class="form-check-label">Tought</label>
						</div>
						<div class="form-check form-check-inline">
							<input class="form-check-input" type="checkbox" v-model="model.advantages" value="undead">
							<label class="form-check-label">Undead</label>
						</div>
					</div>
				</div>
			</div>

			<div v-if="model.id" class="mt-3">
				<div class="row">
					<label class="col-1 col-form-label"></label>
					<h4 class="col-11">{{model.name}} weapons</h4>
					<label class="col-1 col-form-label"></label>
					<div class="col-11">
						<Weapons :model_id="model.id" />
					</div>
				</div>
			</div>
		</div>
		<hr>
	</div>
</template>

<script>
import Weapons from "./weapons.vue";
export default {
	name: "Model",
	props: ["model"],
	components: {
		Weapons
	},
	data() {
		return {
			alert: "",
			alert_succes: false
		};
	},
	methods: {
		save: function(model) {
			if (model.id == null) {
				model.id = 0;
			}
			this.reset();
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/model/" + model.id + "?lang=" + this.$language, model)
				.then(function(res) {
					console.log(res);
					this.alert = "save success";
					this.alert_success = true;
					if (res.status === 201) {
						model.id = res.data;
						this.$emit("add", model);
					}
				})
				.catch(function(err) {
					this.alert = "error: " + err;
					this.alert_success = false;
				});
		},
		remove: function(model) {
			this.reset();
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/model/" + model.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						this.$emit("remove");
					}
				})
				.catch(function(err) {
					this.alert = "error: " + err;
					this.alert_success = false;
				});
		},
		reset: function() {
			this.alert = "";
			this.alert_succes = false;
		},
	}
};
</script>

<style scoped>
.statline input {
	max-width: 4rem;
}
</style>
