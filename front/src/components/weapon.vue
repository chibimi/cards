<template>
	<div class="w-100">
		<div class="form-group row px-3">
			<span class="col-2 text-left">Name</span>
			<span class="col-2 text-left">Type</span>
			<span class="col-1 text-left">rng</span>
			<span class="col-1 text-left">pow</span>
			<span class="col-1 text-left">rof</span>
			<span class="col-1 text-left">aoe</span>
			<span class="col-1 text-left">loc</span>
			<span class="col-1 text-left">cnt</span>
			<span class="col-1 text-left"></span>
			<span class="col-1 text-left"></span>
			
			<input v-model="weapon.name" type="text" class="form-control col-2" placeholder="Name">
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

			<div v-if="weapon.id" class="col-1 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-success" @click="save(weapon)">Update</button>
			</div>
			<div v-if="weapon.id" class="col-1 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-danger" @click="remove(weapon)">Delete</button>
			</div>
			<div v-if="!weapon.id" class="col-2 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-primary" @click="save(weapon)">Add</button>
			</div>

			<div class="col-12 text-left px-0 mt-2">
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="blessed">
					<label class="form-check-label">Blessed</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="chain">
					<label class="form-check-label">Chain</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="type_corrosion">
					<label class="form-check-label">Type: Corrosion</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="continuous_corrosion">
					<label class="form-check-label">Cont. Corrosion</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="crit_corrotion">
					<label class="form-check-label">Crit. Corrosion</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="type_electricity">
					<label class="form-check-label">Type: Electricity</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="disruption">
					<label class="form-check-label">Dusruption</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="crit_disruption">
					<label class="form-check-label">Crit. Disruption</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="type_fire">
					<label class="form-check-label">Type: Fire</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="continuous_fire">
					<label class="form-check-label">Cont. Fire</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="crit_fire">
					<label class="form-check-label">Crit. Fire</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="magical">
					<label class="form-check-label">Magical</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="open_fist">
					<label class="form-check-label">Open Fist</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="shield_1">
					<label class="form-check-label">Shield +1</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="shield_2">
					<label class="form-check-label">Shield +2</label>
				</div>
				<div class="form-check form-check-inline">
					<input class="form-check-input" type="checkbox" v-model="weapon.advantages" value="weapon_master">
					<label class="form-check-label">Weapon Master</label>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "Weapon",
	props: ["weapon"],
	data() {
		return {};
	},
	methods: {
		save: function(weapon) {
			if (weapon.id == null) {
				weapon.id = 0;
			}
			this.$http
				.put("http://localhost:9901/weapons/" + weapon.id, weapon)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						weapon.id = res.data;
						this.$emit("add", weapon);
					}
				});
		},
		remove: function(weapon) {
			this.$http
				.delete("http://localhost:9901/weapons/" + weapon.id)
				.then(function(res) {
					console.log(res);
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

<style scoped>
.statline input {
	max-width: 4rem;
}
</style>
