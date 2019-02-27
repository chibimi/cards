<template>
	<div class="w-100">
		<div class="form-group row">
			<label class="col-2 col-form-label">Model Name 2</label>
			<div class="col-7">
				<input v-model="model.name" type="text" class="form-control" placeholder="Fyanna 2">
			</div>
			<label class="col-1 col-form-label">Order</label>
			<div class="col-2">
				<input v-model="model.order" type="text" class="form-control" placeholder>
			</div>

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
					<input v-model="model.base_size" type="text" class="form-control" placeholder="Base Size">
				</form>
			</div>

			<label class="col-2 col-form-label">Advantages</label>
			<div class="col-10 text-left">
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

			<label class="col-2 col-form-label mt-2"></label>
			<div v-if="model.id" class="col-8 mt-2">
				<button type="submit" class="form-control btn btn-success" @click="save(model)">Update Model</button>
			</div>
			<div v-if="model.id" class="col-2 mt-2">
				<button type="submit" class="form-control btn btn-danger" @click="remove(model.id)">Delete Model</button>
			</div>
			<div v-if="!model.id" class="col-10 mt-2">
				<button type="submit" class="form-control btn btn-primary" @click="save(model)">Save Model</button>
			</div>
		</div>
		<hr>
	</div>
</template>

<script>
export default {
	name: "Model",
	props: ["model"],
	watch: {
		model: function(newVal, oldVal) {
			if (newVal.id === oldVal.id) {
				return;
			}
			console.log("CHANGED MODEL", newVal, oldVal);
			// watch it
			this.get(newVal.id);
		}
	},
	created: function() {
		this.get(this.model.id);
	},
	data() {
		return {};
	},
	methods: {
		save: function(model) {
			if (model.id == null) {
				model.id = 0;
			}
			this.$http
				.put("http://localhost:9901/models/" + model.id, model)
				.then(function(res) {
					console.log(res);
					console.log(res.data);
					if (res.status === 201) {
						model.id = res.data;
						this.$emit("add");
					}
				});
		},
		get: function(id) {
			if (!id) {
				return;
			}
			this.$http
				.get("http://localhost:9901/models/" + id)
				.then(function(res) {
					console.log(res);
					if (res.status === 200) {
						// this.model = res.data;
						this.model.name = res.data.name;
					}
				})
				.catch(function(err) {
					console.log(err);
				});
		},
		remove: function(id) {
			this.$http
				.delete("http://localhost:9901/models/" + id)
				.then(function(res) {
					if (res.status === 204) {
						this.$emit("remove");
					}
				})
				.catch(function(err) {
					console.log(err);
				});
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
