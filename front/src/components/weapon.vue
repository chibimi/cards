<template>
	<div class="w-100">
		<div class="row px-3">
			<div class="col-4">
				<div class="form-group row my-0">
					<span class="col-6 text-left">English Name</span>
					<span class="col-§ text-left">Name</span>
				</div>
			</div>
			<div class="col-8">
				<div class="form-group row my-0">
					<span class="col-2 text-left">Type</span>
					<span class="col-1 text-left">rng</span>
					<span class="col-1 text-left">pow</span>
					<span class="col-1 text-left">rof</span>
					<span class="col-1 text-left">aoe</span>
					<span class="col-1 text-left">loc</span>
					<span class="col-1 text-left">cnt</span>
					<span class="col-1 text-left"></span>
					<span class="col-1 text-left"></span>
				</div>
			</div>

			<div class="col-4">
				<div class="form-group row my-0">
					<input v-model="weapon.title" type="text" class="form-control col-6" placeholder="English Name">
					<input v-model="weapon.name" type="text" class="form-control col-6" placeholder="Name">
				</div>
			</div>
			<div class="col-8">
				<div class="form-group row my-0">
					<select v-model="weapon.type" class="form-control col-2">
						<option value="1">Meele</option>
						<option value="2">Ranged</option>
						<option value="3">Mount</option>
					</select>
					<input v-model="weapon.rng" type="text" class="form-control col-1" placeholder="rng">
					<input v-model="weapon.pow" type="text" class="form-control col-1" placeholder="pow">
					<input v-model="weapon.rof" type="text" class="form-control col-1" placeholder="rof">
					<input v-model="weapon.aoe" type="text" class="form-control col-1" placeholder="aoe">
					<input v-model="weapon.loc" type="text" class="form-control col-1" placeholder="loc">
					<input v-model="weapon.cnt" type="text" class="form-control col-1" placeholder="cnt">

					<div v-if="weapon.id" class="col-2 pl-1 pr-0">
						<button type="submit" class="form-control btn btn-success" @click="save(weapon)">Update</button>
					</div>
					<div v-if="weapon.id" class="col-2 pl-1 pr-0">
						<button type="submit" class="form-control btn btn-danger" @click="remove(weapon)">Delete</button>
					</div>
					<div v-if="!weapon.id" class="col-2 pl-1 pr-0">
						<button type="submit" class="form-control btn btn-primary" @click="save(weapon)">Add</button>
					</div>
				</div>
			</div>

			<div class="col-12 text-left px-0 mt-2">
				<label v-for="a in advantages" :key="a.label" v-bind:value="a.label" class="form-check form-check-inline form-check-label" >
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" :value="a.label">{{a.name}}
				</label>
			</div>
		</div>
	</div>
</template>

<script>
import Tooltip from "./tooltip.vue";
import { WeaponAdvantages } from "./const.js";

export default {
	name: "Weapon",
	props: ["weapon"],
	components: { Tooltip },
	watch: {
		weapon: function(newVal) {
			this.getVO(newVal.id);
		}
	},
	created: function() {
		this.getVO(this.weapon.id);
	},
	data() {
		return {
			vo: {},
			alert: "",
			alert_succes: false,
			advantages: WeaponAdvantages
		};
	},
	methods: {
		getVO: function(id) {
			if (id == null) {
				return;
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/weapon/" + id + "/vo")
				.then(function(res) {
					console.log(res);
					this.vo = res.data;
				});
		},
		save: function(weapon) {
			if (weapon.id == null) {
				weapon.id = 0;
			}
			this.reset();
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/weapon/" + weapon.id + "?lang=" + this.$language, weapon)
				.then(function(res) {
					console.log(res);
					this.alert = "save success";
					this.alert_success = true;
					if (res.status === 201) {
						weapon.id = res.data;
						this.$emit("add", weapon);
					}
				})
				.catch(function(err) {
					this.alert = "error: " + err;
					this.alert_success = false;
				});
		},
		remove: function(weapon) {
			this.reset();
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/weapon/" + weapon.id)
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
